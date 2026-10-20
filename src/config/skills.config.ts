import { BrainCircuit, Crown, Monitor, Server, Cloud, Smartphone } from 'lucide-react';
import type { LucideIcon } from 'lucide-react';

export interface SkillCategoryConfig {
  id: string;
  name: string;
  icon: LucideIcon;
  theme: {
    bg: string;
    text: string;
    pill: string;
    glow: string;
    border: string;
  };
  items: string[];
}

export const SKILLS_CONFIG: SkillCategoryConfig[] = [
  {
    id: 'ai-driven',
    name: 'Agentic AI',
    icon: BrainCircuit,
    theme: {
      bg: 'bg-rose-500/10 dark:bg-rose-500/15',
      text: 'text-rose-700 dark:text-rose-300',
      pill: 'bg-rose-100 dark:bg-rose-900/40 border-rose-200 dark:border-rose-700/50',
      glow: 'hover:shadow-[0_0_12px_rgba(244,63,94,0.15)]',
      border: 'border-rose-500',
    },
    items: [
      'Harness Engineering (Skills, Rules, Hooks, MCP)',
      'Agentic Factories',
      'Context & Memory Engineering',
      'Graph Engineering & Multi-Agent Orchestration',
      'Spec-Driven Development (AGENTS.md, SKILL.md, PRD)',
      'AI Guardrails, Evals & Code Review',
      'Agentic Runtimes (Antigravity, Cursor, Claude Code, Gemini CLI)'
    ],
  },
  {
    id: 'tech-lead',
    name: 'Architecture',
    icon: Crown,
    theme: {
      bg: 'bg-violet-500/10 dark:bg-violet-500/15',
      text: 'text-violet-700 dark:text-violet-300',
      pill: 'bg-violet-100 dark:bg-violet-900/40 border-violet-200 dark:border-violet-700/50',
      glow: 'hover:shadow-[0_0_12px_rgba(139,92,246,0.15)]',
      border: 'border-violet-500',
    },
    items: [
      'System Design (Microservices, Micro-Frontends, BFF)',
      'Domain-Centric (DDD, Clean, Hexagonal)',
      'Event-Driven & CQRS',
      'Monorepo & Feature-Sliced Design (FSD)',
      'API Design & Contracts',
      'Design Principles (SOLID, DRY, KISS, YAGNI)',
      'Technical Roadmap & ADR',
      'Team Leadership & Code Review',
      'Stakeholder Alignment & Agile'
    ],
  },
  {
    id: 'frontend',
    name: 'Frontend',
    icon: Monitor,
    theme: {
      bg: 'bg-amber-500/10 dark:bg-amber-500/15',
      text: 'text-amber-700 dark:text-amber-300',
      pill: 'bg-amber-100 dark:bg-amber-900/40 border-amber-200 dark:border-amber-700/50',
      glow: 'hover:shadow-[0_0_12px_rgba(245,158,11,0.15)]',
      border: 'border-amber-500',
    },
    items: [
      'ECMAScript (ES6+)',
      'TypeScript',
      'React & Next.js (SSR/SSG)',
      'Angular 18+ (Signals)',
      'Zustand & React Query',
      'Micro-Frontends (Module Federation)',
      'Tailwind CSS',
      'HTML5 / CSS3 / SCSS',
      'Web Vitals & Performance',
      'Testing (Jest, Vitest, Playwright)',
      'Legacy Migration (jQuery)'
    ],
  },
  {
    id: 'mobile',
    name: 'Mobile',
    icon: Smartphone,
    theme: {
      bg: 'bg-cyan-500/10 dark:bg-cyan-500/15',
      text: 'text-cyan-700 dark:text-cyan-300',
      pill: 'bg-cyan-100 dark:bg-cyan-900/40 border-cyan-200 dark:border-cyan-700/50',
      glow: 'hover:shadow-[0_0_12px_rgba(6,182,212,0.15)]',
      border: 'border-cyan-500',
    },
    items: [
      'Flutter',
      'Dart',
      'iOS & Android',
      'Riverpod',
      'Local-First Architecture',
      'Offline Sync (SQLite)',
      'Telegram Mini Apps (TMA SDK)',
      'App Store & Google Play Release'
    ],
  },
  {
    id: 'backend',
    name: 'Backend',
    icon: Server,
    theme: {
      bg: 'bg-emerald-500/10 dark:bg-emerald-500/15',
      text: 'text-emerald-700 dark:text-emerald-300',
      pill: 'bg-emerald-100 dark:bg-emerald-900/40 border-emerald-200 dark:border-emerald-700/50',
      glow: 'hover:shadow-[0_0_12px_rgba(16,185,129,0.15)]',
      border: 'border-emerald-500',
    },
    items: [
      'Node.js',
      'NestJS & Express.js',
      'PostgreSQL',
      'Supabase & Firebase',
      'Prisma ORM',
      'Redis',
      'MongoDB',
      'REST API & WebSockets',
      'Auth & Security (OAuth 2.0, JWT, SSO)',
      'Search Engines (OpenSearch, Coveo AI)',
      'Unit & Integration Testing'
    ],
  },
  {
    id: 'devops',
    name: 'DevOps',
    icon: Cloud,
    theme: {
      bg: 'bg-blue-500/10 dark:bg-blue-500/15',
      text: 'text-blue-700 dark:text-blue-300',
      pill: 'bg-blue-100 dark:bg-blue-900/40 border-blue-200 dark:border-blue-700/50',
      glow: 'hover:shadow-[0_0_12px_rgba(59,130,246,0.15)]',
      border: 'border-blue-500',
    },
    items: [
      'Docker',
      'CI/CD (GitHub Actions, Jenkins)',
      'Google Cloud Platform (GCP)',
      'AWS & Oracle Cloud',
      'Cloudflare',
      'Firebase (Hosting, Analytics)',
      'Nginx',
      'Terraform',
      'SonarQube'
    ],
  },
];
