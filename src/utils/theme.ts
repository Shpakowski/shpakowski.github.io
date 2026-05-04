import { SKILLS_CONFIG } from '@/config';

export const getThemeForSkill = (skillName: string) => {
  const lowerSkill = skillName.toLowerCase();
  
  for (let i = 0; i < SKILLS_CONFIG.length; i++) {
    const categorySkills = SKILLS_CONFIG[i].items.map(s => s.toLowerCase());
    
    // Exact match
    if (categorySkills.includes(lowerSkill)) return SKILLS_CONFIG[i].theme;
    
    // Partial match
    for (const catSkill of categorySkills) {
      if (catSkill.length > 2 && (catSkill.includes(lowerSkill) || lowerSkill.includes(catSkill))) {
        return SKILLS_CONFIG[i].theme;
      }
    }
  }
  
  // Fallback dictionary for known terms in experience but not explicitly in skills
  const fallbackMap: Record<string, number> = {
    'go': 2, 'redis': 2, 'typeorm': 2, 'stripe': 2, 'rabbitmq': 2,
    'cloudflare': 3, 'gcp': 3, 'nx': 1, 'vite': 1, 'webpack': 1, 
    'telegram sdk': 1, 'rxjs': 1
  };
  
  if (fallbackMap[lowerSkill] !== undefined) {
    return SKILLS_CONFIG[fallbackMap[lowerSkill]].theme;
  }

  return null; // Fallback to default gray
};
