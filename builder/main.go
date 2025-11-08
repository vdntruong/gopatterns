package main

import "fmt"

func main() {
	// Build a gaming computer
	gamingPC, gamingPCErr := NewComputerBuilder().
		SetCPU("Intel Core i9").
		SetRAM(32).
		SetStorage(1000).
		SetGPU("NVIDIA RTX 4090").
		SetOS("Windows 11").
		Build()
	if gamingPCErr != nil {
		panic(gamingPCErr)
	}

	fmt.Println("Gaming PC:")
	fmt.Println(gamingPC)

	// Build a basic office computer
	officePC, officePCErr := NewComputerBuilder().
		SetCPU("Intel Core i5").
		SetRAM(16).
		SetStorage(512).
		SetOS("Windows 11").
		Build()
	if officePCErr != nil {
		panic(officePCErr)
	}

	fmt.Println("\nOffice PC:")
	fmt.Println(officePC)
}
