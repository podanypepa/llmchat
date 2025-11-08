// Package main demonstrates how to generate an image using the OpenAI API with the llmchat library in Go.
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/podanypepa/llmchat/grok"
	"github.com/podanypepa/llmchat/pkg/imagetools"
)

func main() {
	// Create a new client. For production use, consider a more secure way to manage the API key.
	client, err := grok.NewClient(os.Getenv("GROK_API_KEY"))
	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
		os.Exit(1)
	}

	// Define the image generation request.
	// Using the "grok-2-image-1212" model.
	req := grok.ImageRequest{
		Model:  "grok-2-image-1212",
		Prompt: "a surrealist painting of a clock melting on a tree branch",
		N:      1,
		Size:   "1024x1024",
		// For image generation, it's important to set ResponseFormat to "b64_json".
		ResponseFormat: map[string]string{"type": "b64_json"},
	}

	fmt.Println("Sending image generation request to Grok...")

	// Send the request to the API.
	resp, err := client.SendImageRequest(context.Background(), req)
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		os.Exit(1)
	}

	// Process the response and save the image.
	if len(resp.Data) > 0 && resp.Data[0].B64JSON != "" {
		fmt.Println("Image generated successfully. Saving to disk...")
		err := imagetools.SaveImage(resp.Data[0].B64JSON, "grok_generated_image.png")
		if err != nil {
			fmt.Printf("Error saving image: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Image saved as 'grok_generated_image.png'")
	} else {
		fmt.Println("No image data found in the response.")
	}
}
