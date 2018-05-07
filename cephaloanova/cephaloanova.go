// Package cephaloanova performs analysis of variance on previously grouped data
package cephaloanova

import (
	"math"

	"github.com/paulidealiste/Cephalopod/cephalobjects"
	"github.com/paulidealiste/Cephalopod/cephaloutils"
)

type ssmSummary struct {
	df  float64
	ssm float64
}

type sseSummary struct {
	df  float64
	sse float64
}

func analysisOfVariance(input []cephalobjects.DataPoint) cephalobjects.AnovaSummary {
	var summary cephalobjects.AnovaSummary
	desc := cephaloutils.CalculateDescriptors(input)
	summary.Dft = float64(len(input) - 1)
	summary.SST = desc.VarX * summary.Dft
	channelSSM := make(chan ssmSummary)
	channelSSE := make(chan sseSummary)
	go modelSumOfSquares(input, channelSSM)
	go errorSumOfSquares(input, channelSSE)
	ssm := <-channelSSM
	sse := <-channelSSE
	summary.Dfm = ssm.df
	summary.SSM = ssm.ssm
	summary.Dfe = sse.df
	summary.SSE = sse.sse
	summary.MST = summary.SST / summary.Dft
	summary.MSM = summary.SSM / summary.Dfm
	summary.MSE = summary.SSE / summary.Dfe
	summary.F = summary.MSM / summary.MSE
	return summary
}

func modelSumOfSquares(input []cephalobjects.DataPoint, c chan ssmSummary) {
	var ssmAccumulator, meanAcucumulator, dfAccumulator float64
	var previousDP cephalobjects.DataPoint
	for i, dp := range input {
		meanAcucumulator += dp.X
		if dp.A != previousDP.A || previousDP.A == "" {
			ssmAccumulator += math.Pow(meanAcucumulator/float64(i+1), 2) * float64(i+1)
			dfAccumulator++
			meanAcucumulator = 0.0
		}
		previousDP = dp
	}
	ssm := ssmSummary{df: dfAccumulator - 1, ssm: ssmAccumulator}
	c <- ssm
}

func errorSumOfSquares(input []cephalobjects.DataPoint, c chan sseSummary) {
	var sseAccumulator, dfAccumulator float64
	var previousDP cephalobjects.DataPoint
	var oneGroup []cephalobjects.DataPoint
	for _, dp := range input {
		if dp.A != previousDP.A && previousDP.A != "" {
			groupDesc := cephaloutils.CalculateDescriptors(oneGroup)
			sseAccumulator += groupDesc.VarX * float64(len(oneGroup)-1)
			dfAccumulator += float64(len(oneGroup) - 1)
			oneGroup = nil
		}
		oneGroup = append(oneGroup, dp)
		previousDP = dp
	}
	groupDesc := cephaloutils.CalculateDescriptors(oneGroup)
	sseAccumulator += groupDesc.VarX * float64(len(oneGroup)-1)
	dfAccumulator += float64(len(oneGroup) - 1)
	sse := sseSummary{df: dfAccumulator, sse: sseAccumulator}
	c <- sse
}
