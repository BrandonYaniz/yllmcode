package project

type defaultFile struct {
	Path    string
	Content string
}

var defaultDirectories = []string{
	".yllmcode",
	".yllmcode/decisions",
	".yllmcode/maps",
	".yllmcode/phases",
	".yllmcode/policies",
	".yllmcode/prompts",
	".yllmcode/research",
	".yllmcode/tests",
}

var defaultFiles = []defaultFile{
	{
		Path: ".yllmcode/project.md",
		Content: `# Project

Describe the project, the target users, and the problem it should solve.
`,
	},
	{
		Path: ".yllmcode/instructions.md",
		Content: `# Instructions

Record project-specific development instructions, constraints, and review expectations.
`,
	},
	{
		Path: ".yllmcode/architecture.md",
		Content: `# Architecture

Document the main components, boundaries, data flow, and important tradeoffs.
`,
	},
	{
		Path: ".yllmcode/roadmap.md",
		Content: `# Roadmap

Track implementation phases and the order they should be completed.
`,
	},
	{
		Path: ".yllmcode/dependencies.md",
		Content: `# Dependencies

List required dependencies, why they are needed, and who approved them.
`,
	},
	{
		Path: ".yllmcode/.gitignore",
		Content: `sessions/
state/
scratch/
cache/
logs/
*.tmp
`,
	},
	{
		Path: ".yllmcode/policies/internet-access.md",
		Content: `# Internet Access Policy

Internet access is disabled by default. Record any approved exceptions here.
`,
	},
	{
		Path: ".yllmcode/policies/dependency-policy.md",
		Content: `# Dependency Policy

Prefer the standard library and existing project dependencies. Record approvals for new dependencies here.
`,
	},
}
