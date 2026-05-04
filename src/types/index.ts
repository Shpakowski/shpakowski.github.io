export interface SubProject {
  title: string;
  role: string;
  description: string;
  achievements?: string[];
  stack: string;
}

export interface ExperienceEntry {
  id: string;
  company: string;
  title: string;
  dates: string;
  location?: string;
  role?: string;
  description?: string;
  achievements?: string[];
  stack?: string;
  subProjects?: SubProject[];
}

export interface Highlight {
  metric: string;
  title: string;
  description: string;
}
