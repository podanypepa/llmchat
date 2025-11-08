// Package main demonstrates how to generate an image using the OpenAI API
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/podanypepa/llmchat/gemini"
	"github.com/podanypepa/llmchat/pkg/imagetools"
)

func main() {
	// Create a new client. For production use, consider a more secure way to manage the API key.
	// Note: The model used for image generation is often part of the endpoint URL in Gemini,
	// so the model in the client config might be for chat requests.
	client, err := gemini.NewClient(os.Getenv("GEMINI_API_KEY"), "gemini-pro") // Model is for chat, image model is in the URL.
	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
		os.Exit(1)
	}

	// Define the image generation request.
	req := gemini.ImageRequest{
		Prompt:     "A beautiful landscape painting of a mountain range at sunset",
		ImageCount: 1,
	}

	fmt.Println("Sending image generation request to Gemini...")

	// Send the request to the API.
	resp, err := client.SendImageRequest(context.Background(), req)
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		os.Exit(1)
	}

	// Process the response and save the image.
	if len(resp.Images) > 0 && resp.Images[0].B64JSON != "" {
		fmt.Println("Image generated successfully. Saving to disk...")
		err := imagetools.SaveImage(resp.Images[0].B64JSON, "gemini_generated_image.png")
		if err != nil {
			fmt.Printf("Error saving image: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Image saved as 'gemini_generated_image.png'")
	} else {
		fmt.Println("No image data found in the response.")
	}
}
