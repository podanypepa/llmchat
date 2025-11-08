package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/podanypepa/llmchat/deepseek"
)

func main() {
	// Create a new client. For production use, consider a more secure way to manage the API key.
	client, err := deepseek.NewClient(os.Getenv("DEEPSEEK_API_KEY"))
	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
		os.Exit(1)
	}

	// Define the image generation request.
	// We are using the "Janus-Pro-7B" model and specifying a prompt.
	req := deepseek.ImageRequest{
		Model:  "Janus-Pro-7B",
		Prompt: "a photorealistic image of a futuristic city with flying cars",
		N:      1,
		Size:   "1024x1024",
		// For image generation, it's important to set ResponseFormat to "b64_json".
		ResponseFormat: map[string]string{"type": "b64_json"},
	}

	fmt.Println("Sending image generation request to DeepSeek...")

	// Send the request to the API.
	resp, err := client.SendImageRequest(context.Background(), req)
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		os.Exit(1)
	}

	// Process the response and save the image.
	if len(resp.Data) > 0 && resp.Data[0].B64JSON != "" {
		fmt.Println("Image generated successfully. Saving to disk...")
		err := saveImage(resp.Data[0].B64JSON, "deepseek_generated_image.png")
		if err != nil {
			fmt.Printf("Error saving image: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Image saved as 'deepseek_generated_image.png'")
	} else {
		fmt.Println("No image data found in the response.")
	}
}

// saveImage decodes a base64 string and saves it as an image file.
func saveImage(base64Data, filename string) error {
	data, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return fmt.Errorf("error decoding base64 data: %w", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing file: %w", err)
	}

	return nil
}
