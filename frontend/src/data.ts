/*
 * Static content mirrored from the Go UI package:
 *   internal/ui/menu_view.go  (menuItems)
 *   internal/ui/project_view.go (projectItems)
 *   internal/ui/about_view.go   (aboutLines)
 *
 * This is the authoritative copy of the portfolio text shown in the terminal.
 */

import { type ContentLine, type Project } from './theme'

export const MENU_ENTRIES: readonly { title: string; description: string }[] = [
  { title: 'About', description: 'Hi lol' },
  { title: 'Projects', description: 'What I larp about' },
  { title: 'Contact', description: 'Send me a message' },
  { title: 'PortfolioSSH', description: 'How to SSH in' },
]

export const PROJECTS: readonly Project[] = [
  {
    title: 'Portfolio',
    technology: 'Go, Terraform, GCP, SSH, Docker, Prometheus, Grafana',
    description: 'An overengineered portfolio',
    repoURL: 'https://github.com/CodeKheb/PortfolioSSH',
  },
  {
    title: 'CodeShare',
    technology: 'Node.js, OAuth2, Express.js',
    description: `DevKada Hackathon
CodeShare unifies real-time team chat with live GitHub integration.
Attach a repository to a group chat, and every push, pull request,
and release instantly surfaces as a system message right where your team is already talking.
No notifications to miss. No context to lose. One conversation. Everything you need.`,
    repoURL: 'https://github.com/CodeKheb/HelloWorld-CodeShare',
  },
  {
    title: 'ESPresso',
    technology: 'ESP32, Tauri v2 (Rust), TypeScript, React',
    description: `Tauri 2 + React + TypeScript desktop app for sharing coffee profiles over any
local network. **The WiFi is the DNS** — every device with ESPresso open hosts
a pot, and pots on the same network find each other automatically via mDNS.`,
    repoURL: 'https://github.com/CodeKheb/ESPresso',
  },
  {
    title: 'typetest_TUI',
    technology: 'Go, Bubbletea, Lipgloss',
    description: `A terminal-based typing speed test written in Go, using
[Bubbletea](https://github.com/charmbracelet/bubbletea) and
[Lipgloss](https://github.com/charmbracelet/lipgloss) for the UI.`,
    repoURL: 'https://github.com/CodeKheb/typetest_TUI',
  },
  {
    title: 'OrderUp',
    technology: 'JavaFX, FXGL',
    description: `A 2D restaurant management game built with FXGL (JavaFX)
that visualizes CPU scheduling algorithms through a restaurant.
Customers arrive as processes with Arrival Time (AT) and Burst Time (BT),
simulating a First Come First Serve (FCFS) scheduling algorithm.
The game also includes a rhythm minigame where players click to serve customers at the right time.`,
    repoURL: 'https://github.com/CodeKheb/OrderUp',
  },
  {
    title: 'dotfiles',
    technology: 'Neovim, Linux, Lua, Bash',
    description: `Personal Linux configuration files,
built around an Arch-based setup with Neovim as the primary editor,
a Sway (Wayland) desktop, and a set of rofi-powered launcher scripts.`,
    repoURL: 'https://github.com/CodeKheb/dotfiles',
  },
  {
    title: 'Clingy',
    technology: 'TensorFlow, React-Native, Google OAuth2',
    description: `RAITE 2026 Hackathon
A Clingy reminder to do your tasks.
Your Google Classroom deadlines, turned into a study plan.
An Android app that works offline, with a small star called Cling that nudges you.`,
    repoURL: 'https://github.com/CodeKheb/Clingy',
  },
]

// aboutLines, ported from internal/ui/about_view.go
export const ABOUT_LINES: readonly ContentLine[] = [
  // Identity
  { text: ' ', style: 'body' },
  { text: ' ', style: 'body' },
  { text: 'Kherbin Clloyde Buenaventura', style: 'name' },
  { text: 'Age: \t\t20 years old', style: 'muted' },
  { text: 'Location:    Based in Bataan, Philippines', style: 'muted' },
  { text: ' ', style: 'body' },
  { text: ' ', style: 'body' },

  { text: 'Computer Science Student @ Bataan Peninsula State University', style: 'body' },
  { text: 'Vice President for Externals - Student Society of Information Technology Education', style: 'body' },
  { text: ' ', style: 'body' },
  { text: ' ', style: 'body' },

  // Introduction
  { text: 'What this project is:', style: 'section' },
  {
    text: 'This is an SSH portfolio written in the Go programming language and deployed in a Google VM.',
    style: 'body',
  },
  { text: ' ', style: 'body' },

  { text: 'A question you might ask is "Why SSH?"', style: 'body' },
  { text: ' ', style: 'body' },

  {
    text: 'Instead of making a website portfolio and deploying it on a platform that does everything for me,',
    style: 'body',
  },
  { text: 'I wanted to learn how deployment works internally.', style: 'body' },
  { text: ' ', style: 'body' },

  { text: 'Learning things like:', style: 'body' },
  { text: '    • Terraform', style: 'body' },
  { text: '    • Google Cloud', style: 'body' },
  { text: '    • Networking', style: 'body' },
  { text: '    • SSH', style: 'body' },
  { text: ' ', style: 'body' },

  { text: 'It is also a good insight into my terminal-centric workflow.', style: 'body' },
  { text: ' ', style: 'body' },

  {
    text: 'I am an aspiring DevOps engineer, and this is also why I made a terminal portfolio.',
    style: 'body',
  },
  { text: ' ', style: 'body' },

  // Education
  { text: 'EDUCATION', style: 'section' },
  { text: ' ', style: 'body' },

  { text: 'Bachelor of Science in Computer Science', style: 'subheading' },
  { text: 'Bataan Peninsula State University | 2025 - Present', style: 'muted' },
  { text: '    • President\'s Lister', style: 'body' },
  { text: '    • GWA (1.3x)', style: 'body' },
  { text: ' ', style: 'body' },

  {
    text: 'Technical Vocational and Livelihood - Information and Communications Technology',
    style: 'subheading',
  },
  { text: 'Bataan National High School | 2023 - 2025', style: 'muted' },
  { text: '    • With High Honors', style: 'body' },
  {
    text: '    • Awarded Outstanding in Computer Systems Servicing (Grade: 100%)',
    style: 'body',
  },
  { text: ' ', style: 'body' },

  // Experiences
  { text: 'EXPERIENCES', style: 'section' },
  { text: ' ', style: 'body' },

  {
    text: 'Regional Assembly on Information Technology Education (RAITE)',
    style: 'subheading',
  },
  { text: 'Quiz Bee Representative | 2025', style: 'muted' },
  { text: 'Nueva Ecija University of Science and Technology', style: 'muted' },
  {
    text: '    • Competed as the Quiz Bee Representative for Bataan Peninsula State University',
    style: 'body',
  },
  { text: ' ', style: 'body' },

  {
    text: 'Student Society of Information Technology Education',
    style: 'subheading',
  },
  { text: 'Vice Chairperson for External Affairs | 2025', style: 'muted' },
  {
    text: '    • Helped with the first external sponsorship of SSITE',
    style: 'body',
  },
  { text: ' ', style: 'body' },

  { text: 'National Service Training Program (NSTP)', style: 'subheading' },
  { text: 'Project Lead | 2026', style: 'muted' },
  { text: 'Leadership Awardee | 2026', style: 'muted' },
  {
    text: '    • Led the Literacy Training Service cluster to complete a Campus Based Project',
    style: 'body',
  },
  {
    text: '      under significant budget and time constraints.',
    style: 'body',
  },
  { text: ' ', style: 'body' },

  // Skillset & Languages
  { text: 'SKILLSET & LANGUAGES', style: 'section' },
  { text: ' ', style: 'body' },

  { text: 'DevOps and Infrastructure', style: 'subheading' },
  { text: '    • Docker', style: 'body' },
  { text: '    • Terraform', style: 'body' },
  { text: '    • GitHub Actions', style: 'body' },
  { text: '    • Linux', style: 'body' },
  { text: '    • Prometheus', style: 'body' },
  { text: '    • Grafana', style: 'body' },
  { text: '    • Google Cloud Platform (GCP)', style: 'body' },
  { text: ' ', style: 'body' },

  { text: 'Backend and Database', style: 'subheading' },
  { text: '    • Go', style: 'body' },
  { text: '    • Node.js', style: 'body' },
  { text: '    • Java', style: 'body' },
  { text: '    • C#', style: 'body' },
  { text: '    • PostgreSQL', style: 'body' },
  { text: '    • SQLite', style: 'body' },
  { text: ' ', style: 'body' },

  { text: 'Systems Programming and Embedded', style: 'subheading' },
  { text: '    • C', style: 'body' },
  { text: '    • Rust', style: 'body' },
  { text: '    • Arduino', style: 'body' },
  { text: '    • ESP32', style: 'body' },
  { text: '    • ESP8266', style: 'body' },
  { text: ' ', style: 'body' },

  { text: 'Scripting and Automation', style: 'subheading' },
  { text: '    • Bash', style: 'body' },
  { text: '    • Lua', style: 'body' },
  { text: '    • Python', style: 'body' },
  { text: '    • Maven', style: 'body' },
  { text: '    • Gradle', style: 'body' },
  { text: ' ', style: 'body' },

  { text: 'Frontend and Frameworks', style: 'subheading' },
  { text: '    • Tauri v2 (Rust)', style: 'body' },
  { text: '    • Express.js', style: 'body' },
  { text: '    • React.js', style: 'body' },
  { text: '    • Tailwind CSS', style: 'body' },
  { text: '    • Alpine.js', style: 'body' },
  { text: '    • Dart', style: 'body' },
  { text: '    • Flutter', style: 'body' },
  { text: ' ', style: 'body' },

  { text: 'Tools for Development', style: 'subheading' },
  { text: '    • Arch Linux', style: 'body' },
  { text: '    • Neovim', style: 'body' },
  { text: '    • Blender', style: 'body' },
  { text: '    • Godot', style: 'body' },
  { text: '    • Git', style: 'body' },
  { text: ' ', style: 'body' },

  // Certifications
  { text: 'CERTIFICATIONS', style: 'section' },
  { text: ' ', style: 'body' },

  { text: 'IC3 Digital Literacy Certification - Level 1', style: 'subheading' },
  { text: 'October 2025', style: 'muted' },
  { text: ' ', style: 'body' },

  { text: 'ITS: Java Certification', style: 'subheading' },
  { text: 'April 2026', style: 'muted' },
  { text: ' ', style: 'body' },
  { text: ' ', style: 'body' },

  { text: 'IC3 Digital Literacy Certification - Level 2', style: 'subheading' },
  { text: 'October 2026', style: 'muted' },
  { text: ' ', style: 'body' },
]

// Header titles per screen, ported from the WindowSizeMsg handler in model.go
export function headerTitle(
  screen: 'menu' | 'about' | 'projects' | 'readme' | 'contact' | 'ssh',
  projectTitle?: string,
): string {
  switch (screen) {
    case 'menu':
      return "Kherbin's Portfolio"
    case 'about':
      return 'ABOUT KHERBIN BUENAVENTURA'
    case 'projects':
      return 'PROJECTS'
    case 'readme':
      return projectTitle ? `PROJECT: ${projectTitle}` : 'PROJECT'
    case 'contact':
      return 'CONTACT'
    case 'ssh':
      return 'PORTFOLIO SSH'
    default:
      return "Kherbin's Portfolio"
  }
}

// The text shown when a project has no README, ported from readme_view.go
export const NO_README_TEXT = "I'm lazy and did not write a README for this yet"
