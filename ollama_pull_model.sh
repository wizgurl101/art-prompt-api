#!/bin/bash

/bin/ollama serve &
pid=$!

sleep 5

echo "🔴 Retrieve LLAMA3.2 1b model..."
ollama pull llama3.2:1b
echo "🟢 Finish pulling model."

wait $pid