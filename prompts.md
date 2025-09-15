
# prompts.md

This document contains some of the prompts I used with GenAI tools during the development of this project.

## Prompt 1 – Backend PRD
```
## Requirements

### Backend: API Development

* Implement an API that supports the frontend by providing the necessary product details
* The primary endpoint should fetch product details

### Stack:

* You can use any backend technology or framework of your choice
* Do not use real databases, persist everything in local JSON or CSV files

### Non-functional requirements:
Special consideration will be given to good practices in error handling, documentation, testing, and any other relevant non-functional aspects you choose to demonstrate

### Tool Usage
Allowed Tools: You may use and are encouraged to use GenAI tools, agentic IDEs, and other code assistance tools to help generate ideas or code

### Documentation & strategic overview
Please include a brief README o Diagram (optional) that explains your API design, main endpoints, setup instructions, and any key architectural decisions you made during development

### Technical strategy:
Detail the chosen technology stack for backend
Explain how GenAI and modern development tools are integrated to improve efficiency

### Submission
It must contain a run.md explaining how to run the project
In case any AI productivity tool was used, it is greatly appreciated if you can share the different prompts that were used on *prompts.md* inside the project

---

Help me generate a PRD file for these requirements, using the example structure I provided.
```

## Prompt 2 – Frontend PRD (Next.js 14 App Router)
```
Can you also create a PRD for the frontend of this backend?

I plan to build it using Next.js 14 with the App Router.
```

## Prompt 3 – README and run.md
```
In English, please, and can you also generate README.md and run.md versions?
```

## Prompt 4 – Extra Feature (Filters)
```
Can you include in both PRDs filters by name, price range, and category?
```

## Tool Usage

TaskMaster AI → Generated actionable tasks directly from the backend and frontend PRDs, helping structure the roadmap into implementable steps.

Vercel v0 → Used the frontend PRD as input to scaffold a marketplace-style UI in Next.js, providing a strong starting point for the design and layout.

CodiumAI / Test Generation → Using the Go backend code for the product API, generate unit tests for the following functions:
- Fetch all products
- Fetch product by ID
- Apply filters by name, price range, and category

Include edge cases and error handling tests.

