import { SKILLS_CONFIG } from '@/config';

const getCategoryTheme = (id: string) =>
  SKILLS_CONFIG.find((c) => c.id === id)?.theme ?? null;

export const getThemeForSkill = (skillName: string) => {
  const lowerSkill = skillName.toLowerCase().trim();

  // Explicit category mapping for composite or shorthand stack labels
  const explicitCategoryMap: Record<string, string> = {
    // Agentic AI ('ai-driven')
    'harness & context engineering': 'ai-driven',
    'harness engineering': 'ai-driven',
    'context engineering': 'ai-driven',
    'agentic factories': 'ai-driven',
    'spec-driven development': 'ai-driven',
    'mcp': 'ai-driven',
    'model context protocol': 'ai-driven',
    'prompt engineering': 'ai-driven',
    'agent orchestration': 'ai-driven',
    'coveoui ai': 'ai-driven',

    // Architecture ('tech-lead')
    'system design': 'tech-lead',
    'ddd & clean architecture': 'tech-lead',
    'ddd': 'tech-lead',
    'clean architecture': 'tech-lead',
    'hexagonal architecture': 'tech-lead',
    'monorepo': 'tech-lead',
    'fsd': 'tech-lead',
    'agile': 'tech-lead',

    // Frontend ('frontend')
    'react': 'frontend',
    'next.js': 'frontend',
    'next.js & react (ssr/ssg)': 'frontend',
    'angular': 'frontend',
    'angularjs': 'frontend',
    'rxjs': 'frontend',
    'redux': 'frontend',
    'ngrx': 'frontend',
    'zustand': 'frontend',
    'zod': 'frontend',
    'heroui': 'frontend',
    'module federation': 'frontend',
    'bem': 'frontend',
    'sass': 'frontend',
    'html5/css3': 'frontend',
    'javascript (jquery)': 'frontend',
    'wordpress': 'frontend',
    'twig': 'frontend',
    'jasmine': 'frontend',
    'karma': 'frontend',
    'adobe cq5': 'frontend',
    'adobe aem': 'frontend',
    'nx': 'frontend',
    'vite': 'frontend',
    'webpack': 'frontend',

    // Mobile ('mobile')
    'flutter': 'mobile',
    'dart': 'mobile',
    'riverpod': 'mobile',
    'sqlite': 'mobile',
    'telegram mini apps (tma sdk)': 'mobile',
    'telegram sdk': 'mobile',

    // Backend ('backend')
    'node.js': 'backend',
    'nestjs': 'backend',
    'express.js': 'backend',
    'postgresql': 'backend',
    'postgis': 'backend',
    'pgvector': 'backend',
    'prisma': 'backend',
    'prisma orm': 'backend',
    'redis': 'backend',
    'supabase': 'backend',
    'firebase': 'backend',
    'websockets': 'backend',
    'oauth 2.0': 'backend',
    'rest api': 'backend',
    'soap': 'backend',
    'opensearch': 'backend',

    // DevOps ('devops')
    'docker': 'devops',
    'nginx': 'devops',
    'github actions': 'devops',
    'jenkins': 'devops',
    'gcp': 'devops',
    'google cloud platform (gcp)': 'devops',
    'oracle cloud': 'devops',
    'cloudflare': 'devops',
    'sonarqube': 'devops',
    'git': 'devops',
  };

  if (explicitCategoryMap[lowerSkill]) {
    return getCategoryTheme(explicitCategoryMap[lowerSkill]);
  }

  for (let i = 0; i < SKILLS_CONFIG.length; i++) {
    const categorySkills = SKILLS_CONFIG[i].items.map((s) => s.toLowerCase());

    // Exact match
    if (categorySkills.includes(lowerSkill)) return SKILLS_CONFIG[i].theme;

    // Partial match
    for (const catSkill of categorySkills) {
      if (catSkill.length > 2 && (catSkill.includes(lowerSkill) || lowerSkill.includes(catSkill))) {
        return SKILLS_CONFIG[i].theme;
      }
    }
  }

  return null; // Fallback to default gray
};
