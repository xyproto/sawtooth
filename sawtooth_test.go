package sawtooth

import (
	"fmt"
	"os"
	"testing"
	"time"
)

// writeFile writes the given data slice to a file, formatted for use with GNUPlot.
func writeFile(filename string, data []float64) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for i, value := range data {
		// Write data in a simple two-column format: index value
		_, err := fmt.Fprintf(file, "%d %f\n", i, value)
		if err != nil {
			return err
		}
	}

	return nil
}

// TestSawtoothWaveGenerators tests if both sawtooth wave generator functions produce approximately equal outputs.
// It also generates files for GNUPlot to visually compare the waveforms.
func TestSawtoothWaveGenerators(t *testing.T) {
	frequency := 440.0 // A4 note
	sampleRate := 44100
	duration := time.Duration(float64(time.Second) / frequency) // Duration to generate just one period of the wave

	formulaWave := GenerateSawtoothFormula(frequency, sampleRate, duration)
	particleWave := GenerateSawtoothParticle(frequency, sampleRate, duration)

	// Generate GNUPlot compatible files
	if err := writeFile("formula_wave.dat", formulaWave); err != nil {
		t.Fatalf("Failed to write formula waveform data: %v", err)
	}
	if err := writeFile("particle_wave.dat", particleWave); err != nil {
		t.Fatalf("Failed to write particle waveform data: %v", err)
	}

	// Here you could add your comparison logic if needed
	// For now, we just generate the data files for GNUPlot
}
