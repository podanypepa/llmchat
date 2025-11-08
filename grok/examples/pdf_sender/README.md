# PDF Sender Example

This example demonstrates how to send the content of a PDF file to the Grok API for summarization.

## Dependencies

This example requires the `rsc.io/pdf` package to extract text from PDF files.

Install it using:
```sh
go get rsc.io/pdf
```

## How to Run

1.  **Set your X-AI API key:**
    ```sh
    export XAI_API_KEY="your-api-key"
    ```

2.  **Run the application with the path to your PDF file:**
    ```sh
    go run main.go /path/to/your/document.pdf
    ```

The application will extract the text from the PDF, send it to the Grok API, and print the summary.
