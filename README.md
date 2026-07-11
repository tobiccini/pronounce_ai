AI-powered pronunciation assistant for broadcasters, journalists and news presenters.

# Pronounce AI

> **AI-powered pronunciation assistant for broadcasters, journalists, and news presenters.**

---

## Overview

Pronounce AI helps broadcasters confidently pronounce unfamiliar names, places, organizations, and foreign terms before going live.

Simply paste a news script into the application, and the AI analyzes it in seconds, identifying words that may be difficult to pronounce while providing clear pronunciation guidance, phonetic transcriptions, contextual meanings, presenter tips, and confidence scores.

The project was built to improve newsroom efficiency, reduce pronunciation errors during live broadcasts, and make international news more accessible to presenters.

---

# Why Pronounce AI?

Newsrooms regularly cover international stories involving names from different languages and cultures.

Examples include:

* Volodymyr Zelenskyy
* Mark Rutte
* Alassane Ouattara
* António Guterres
* Xi Jinping
* Oleksandr Usyk

Finding the correct pronunciation often requires searching multiple websites or videos before a live broadcast.

Pronounce AI automates this process.

---

# Features

## AI-powered pronunciation analysis

Analyzes complete news scripts using Large Language Models hosted on Fireworks AI.

For every detected difficult word, the application returns:

* Easy pronunciation guide
* English IPA
* Native IPA
* Language of origin
* Meaning
* Presenter tips
* Difficulty rating
* Confidence score
* Suggested broadcast replacement

---

## Intelligent analysis caching

Every analyzed script is hashed using SHA-256.

If the exact same script is analyzed again, results are instantly retrieved from PocketBase instead of calling the AI model again.

Benefits:

* Faster response time
* Reduced AI costs
* Better user experience
* Lower API usage

---

## Automatic AI fallback

Reliability is essential during live demonstrations and production use.

Pronounce AI automatically retries analysis if the primary AI model becomes unavailable.

Primary model:

Nemotron 3 Ultra (nvfp4)

↓

Fallback:

DeepSeek V4 Flash

This happens transparently without interrupting the user.

---

## Local pronunciation dictionary

Broadcasters can maintain preferred newsroom pronunciations.

Dictionary entries override AI-generated results, ensuring consistent pronunciation across broadcasts.

---

## Analysis history

Every successful analysis is stored in PocketBase for future reference.

---

## Processing time

Each analysis displays its execution time, allowing users to compare cached responses with fresh AI inference.

---

# Architecture

```
                    Flutter Application
                           │
                    REST API (JSON)
                           │
                           ▼
                    Go Backend Server
               ┌───────────┴───────────┐
               │                       │
               ▼                       ▼
        PocketBase Database      Fireworks AI
      (Cache • History • DB)   (LLM Inference)
                                       │
                                       ▼
                             DeepSeek V4 Pro
                                       │
                              Automatic Fallback
                                       ▼
                             DeepSeek V4 Flash
```

---

# AI Workflow

```
Paste News Script
        │
        ▼
Generate SHA-256 Hash
        │
        ▼
Script already analyzed?
     │               │
    Yes              No
     │               │
     ▼               ▼
Return Cached    Fireworks AI
Analysis             │
                     ▼
             DeepSeek V4 Pro
                     │
             Successful?
              │          │
             Yes         No
              │          ▼
              │   DeepSeek V4 Flash
              │          │
              └──────────┘
                     │
                     ▼
      Apply Local Dictionary Overrides
                     │
                     ▼
 Store Analysis + Pronunciations
                     │
                     ▼
       Return Results to Flutter
```

---

# Technology Stack

## Frontend

* Flutter

## Backend

* Go

## Local Database

* PocketBase

## AI Platform

* Fireworks AI

## Language Models

* Nemotron 3 Ultra (nvfp4)
* DeepSeek V4 Flash

---

# Project Structure

```
pronounce-ai/

├── app/
│   ├── lib/
│   ├── assets/
│   └── pubspec.yaml
│

│
├── pocketbase/
│   ├── pb_data/
│   └── pb_hooks/
        └── ai/
│             └── builder.go
              └── dictionary.go
              └── fireworks.go
              └── gemini.go
              └── indexer.go
              └── models.go
              └── prompts.go
              └── provider.go
              └── storage.go
              └── main.go
        └── api
              └── pronounce.go    

└── docs   
│   ├── architecture.png
│   ├── ai-workflow.png
│   ├── demo-script.md
│   └── screenshots/
│
├── README.md
├── LICENSE
├── .gitignore
└── .env.example
```

---

# Screenshots

Add screenshots to:

```
docs/screenshots/
```

Suggested screenshots:

* Home Screen
* Analysis Results
* Bottom Sheet
* Analysis History
* Cache Demonstration
* Loading Animation

---

# Installation

## Clone the repository

```bash
git clone <repository-url>
```

---

## Flutter

```bash
cd app

flutter pub get

flutter run
```

---

## Backend

```bash
cd backend

go run . serve
```

---


# Environment Variables

Create a `.env` file inside the backend directory.

```
FIREWORKS_API_KEY=YOUR_API_KEY

FIREWORKS_MODEL=accounts/fireworks/models/deepseek-v4-flash
```

---

# Future Roadmap

* Hybrid entity-level pronunciation cache
* Offline pronunciation support
* Audio pronunciation playback
* Shared newsroom dictionaries
* Team collaboration
* Multi-language newsroom support
* Pronunciation analytics dashboard

---

# Engineering Highlights

* Script-level caching using SHA-256 hashing
* Persistent storage with PocketBase
* Automatic AI model fallback
* Provider-independent backend architecture
* Configurable AI model selection via environment variables
* Clean separation between Flutter, Go, and database layers

---

# Why PocketBase?

PocketBase provides a lightweight backend for:

* Cached analyses
* Pronunciation records
* Analysis history
* Dictionary overrides

Its simplicity and embedded database make it well suited for rapid development and deployment.

---

# Why Fireworks AI?

Fireworks AI provides access to high-performance language models with low-latency inference.

The backend supports configurable models and automatic fallback, allowing the application to remain operational if the primary model is unavailable.

---

# Built For

AMD Developer Hackathon — Track 3: Unicorn (Open Innovation)

---

# License

This project is licensed under the MIT License.

See the LICENSE file for details.
