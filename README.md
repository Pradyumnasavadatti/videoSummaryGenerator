# Video Summarizer CLI 🎥📝

A command-line application built in **Go** using **Cobra** that generates concise summaries from videos.

The tool supports:

- 📁 **Local video files**
- ▶️ **YouTube videos**

It extracts audio from the video, transcribes it using **OpenAI Whisper**, and then generates a summary using the **ChatGPT (GPT-3.5 Turbo) API**.  
The final summary is saved to a file called `summary.txt`.

---

## 🚀 Features

- CLI-based workflow using Cobra
- Supports both local videos and YouTube links
- Automatic audio extraction using FFmpeg
- Speech-to-text using OpenAI Whisper
- AI-generated summaries using ChatGPT
- Saves output summary to `summary.txt`

---

## 🛠 Tech Stack

- **Language**: Go
- **CLI Framework**: Cobra
- **Audio Processing**: FFmpeg
- **Speech-to-Text**: OpenAI Whisper API
- **Summarization**: OpenAI ChatGPT (GPT-3.5 Turbo)

---

## 📦 Prerequisites

Make sure you have the following installed:

- Go `>= 1.20`
- FFmpeg
- An OpenAI API key

Set your OpenAI API key as an environment variable:

```bash
export OPEN_AI_KEY=your_api_key_here
```
