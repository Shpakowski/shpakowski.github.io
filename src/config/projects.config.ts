export interface ProjectConfig {
  id: string;
  nameKey: string;
  descriptionKey: string;
  /** Path relative to /public */
  logo: string;
  /** Accent color for top-bar and hover glow */
  accent: {
    bar: string;       // Tailwind bg class
    glow: string;      // box-shadow rgba string
    border: string;    // Tailwind border-color class on hover
  };
  href: string;
  stack?: string[];
}

export const PROJECTS_CONFIG: ProjectConfig[] = [
  {
    id: 'mcp-coop-dao',
    nameKey: 'projects.mcpCoopDao.name',
    descriptionKey: 'projects.mcpCoopDao.description',
    logo: '/logos/mcp-coop-logo.png',
    accent: {
      bar: 'bg-blue-500',
      glow: 'rgba(59,130,246,0.12)',
      border: 'hover:border-blue-400/40',
    },
    href: '',
    stack: [
      'TypeScript',
      'Next.js & React (SSR/SSG)',
      'Tailwind CSS',
      'Zustand',
      'NestJS',
      'PostgreSQL',
      'OAuth 2.0',
      'DDD',
      'Hexagonal Architecture',
      'Google Cloud Platform (GCP)'
    ],
  },
  {
    id: 'unbogi',
    nameKey: 'projects.unbogi.name',
    descriptionKey: 'projects.unbogi.description',
    logo: '/logos/unbogi-icon.png',
    accent: {
      bar: 'bg-violet-500',
      glow: 'rgba(139,92,246,0.12)',
      border: 'hover:border-violet-400/40',
    },
    href: '',
    stack: [
      'Prompt Engineering',
      'Agent Orchestration',
      'React Native',
      'React',
      'Zod',
      'Firebase',
      'Monorepo'
    ],
  },
  {
    id: 'task-manager',
    nameKey: 'projects.taskManager.name',
    descriptionKey: 'projects.taskManager.description',
    logo: '/logos/mcp-kanban-logo.png',
    accent: {
      bar: 'bg-emerald-500',
      glow: 'rgba(16,185,129,0.12)',
      border: 'hover:border-emerald-400/40',
    },
    href: '',
    stack: [
      'React',
      'Zustand',
      'Module Federation',
      'Zod',
      'HeroUI',
      'NestJS',
      'PostgreSQL',
      'Prisma ORM',
      'Docker',
      'Nginx',
      'GitHub Actions',
      'FSD'
    ],
  },
];
