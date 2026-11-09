// Package spentenergy предоставляет функции для расчета дистанции,
// средней скорости и сожженных калорий при беге и ходьбе
package spentenergy

import (
	"errors"
	"time"
)

const (
	mInKm                      = 1000
	minInH                     = 60
	stepLengthCoefficient      = 0.45
	walkingCaloriesCoefficient = 0.5
)

// ErrInvalidInput возвращается, когда переданы некорректные параметры
var ErrInvalidInput = errors.New("некорректные входные параметры")

// WalkingSpentCalories рассчитывает количество калорий, сожженных при ходьбе
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	runningCalories, err := RunningSpentCalories(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}

	return runningCalories * walkingCaloriesCoefficient, nil
}

// RunningSpentCalories рассчитывает количество калорий, сожженных при беге
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, ErrInvalidInput
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMin := duration.Minutes()

	return (durationInMin * meanSpeed * weight) / minInH, nil
}

// MeanSpeed рассчитывает среднюю скорость движения (км/ч)
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 || height <= 0 || duration <= 0 {
		return 0
	}

	distance := Distance(steps, height)

	return distance / duration.Hours()
}

// Distance рассчитывает пройденную дистанцию в километрах
func Distance(steps int, height float64) float64 {
	if steps <= 0 || height <= 0 {
		return 0
	}

	stepLength := height * stepLengthCoefficient
	distanceInM := stepLength * float64(steps)

	return distanceInM / mInKm
}
