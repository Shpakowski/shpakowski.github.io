import { BrainCircuit, Crown, Monitor, Server, Cloud } from 'lucide-react';
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
    name: 'AI-Driven',
    icon: BrainCircuit,
    theme: {
      bg: 'bg-rose-500/10 dark:bg-rose-500/15',
      text: 'text-rose-700 dark:text-rose-300',
      pill: 'bg-rose-100 dark:bg-rose-900/40 border-rose-200 dark:border-rose-700/50',
      glow: 'hover:shadow-[0_0_12px_rgba(244,63,94,0.15)]',
      border: 'border-rose-500',
    },
    items: [
      'Antigravity',
      'Cursor',
      'Prompt Engineering',
      'Agentic Workflows & Orchestration',
      'MCP',
      'RAG',
      'Anthropic Patterns',
      'Gemini CLI',
      'Automated Code Review & Validation'
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
      'TypeScript',
      'Angular + RxJS',
      'Next.js & React (SSR/SSG)',
      'React Native',
      'Zustand',
      'NgRx',
      'TanStack Query',
      'HTML5/CSS3',
      'Tailwind CSS',
      'SASS/SCSS',
      'BEM',
      'Module Federation (Microfrontends)',
      'Web Vitals & Performance Optimization',
      'Unit Testing',
      'JavaScript (jQuery)'
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
      'NestJS',
      'Express.js',
      'PostgreSQL',
      'Prisma ORM',
      'MongoDB',
      'Firebase',
      'OAuth 2.0',
      'CoveoUI AI',
      'Unit Testing',
      'Integration Testing'
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
      'Google Cloud Platform (GCP)',
      'Oracle Cloud',
      'AWS',
      'Terraform',
      'Docker',
      'Nginx',
      'Kubernetes',
      'GitHub Actions',
      'SonarQube'
    ],
  },
  {
    id: 'tech-lead',
    name: 'Tech Lead',
    icon: Crown,
    theme: {
      bg: 'bg-violet-500/10 dark:bg-violet-500/15',
      text: 'text-violet-700 dark:text-violet-300',
      pill: 'bg-violet-100 dark:bg-violet-900/40 border-violet-200 dark:border-violet-700/50',
      glow: 'hover:shadow-[0_0_12px_rgba(139,92,246,0.15)]',
      border: 'border-violet-500',
    },
    items: [
      'System Design & Architecture Patterns',
      'DDD',
      'Event-Driven',
      'CQRS',
      'FSD',
      'Hexagonal',
      'DRY',
      'SOLID',
      'KISS',
      'Microservices & Microfrontend Architecture',
      'API Design & Contracts',
      'Technical Roadmap & Decision Records (ADR)',
      'Team Leadership (2–5 devs)',
      'Code Review Culture',
      'Stakeholder Communication',
      'Technical Grooming',
      'Agile'
    ],
  },
];
