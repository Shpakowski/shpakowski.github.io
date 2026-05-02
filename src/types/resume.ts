export interface Contact {
  type: string;
  value: string;
  url?: string;
}

export interface Wallet {
  type: string;
  address: string;
}

export interface Skill {
  id: string;
  name: string;
  category?: string;
  level?: number; // 1-5 or similar if needed
}

export interface Experience {
  id: string;
  companyName: string;
  projectRole: string;
  startDate: string; // ISO or formatted
  endDate: string | null; // null means present
  description: string;
  skills?: Skill[];
  location?: string;
}

export interface Education {
  id: string;
  institution: string;
  degree: string;
  startDate: string;
  endDate: string;
  description?: string;
}

export interface Profile {
  id: string;
  fullName: string;
  headline: string;
  email: string;
  location: string;
  timezone: string;
  availabilityStatus?: string;
  languages: string[];
  avatarUrl?: string;
  shortDescription?: string;
  description?: string;
  contacts: Record<string, Contact>;
  wallets?: Record<string, Wallet>;
  skills: Skill[];
  experiences: Experience[];
  educations: Education[];
}
