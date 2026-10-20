export interface ProjectConfig {
  id: string;
  nameKey: string;
  descriptionKey: string;
  /** Path relative to /public or external URL */
  logo: string;
  /** Accent color for top-bar and hover glow */
  accent: {
    bar: string;       // Tailwind bg class
    glow: string;      // box-shadow rgba string
    border: string;    // Tailwind border-color class on hover
  };
  href?: string;
  iosHref?: string;
  androidHref?: string;
  stack?: string[];
}

export const PROJECTS_CONFIG: ProjectConfig[] = [
  {
    id: 'holdy',
    nameKey: 'projects.holdy.name',
    descriptionKey: 'projects.holdy.description',
    logo: 'https://mcpcoop.org/holdy_icon.webp',
    accent: {
      bar: 'bg-blue-500',
      glow: 'rgba(59,130,246,0.12)',
      border: 'hover:border-blue-400/40',
    },
    href: 'https://holdy-app.web.app/download',
  },
  {
    id: 'beerhandshake',
    nameKey: 'projects.beerhandshake.name',
    descriptionKey: 'projects.beerhandshake.description',
    logo: 'https://mcpcoop.org/beerhandshake_icon.webp',
    accent: {
      bar: 'bg-amber-500',
      glow: 'rgba(245,158,11,0.12)',
      border: 'hover:border-amber-400/40',
    },
    iosHref: 'https://apps.apple.com/us/app/beerhandshake/id6772492609',
    androidHref: 'https://play.google.com/store/apps/details?id=org.mcpcoop.beer.handshake',
  },
  {
    id: 'znichka',
    nameKey: 'projects.znichka.name',
    descriptionKey: 'projects.znichka.description',
    logo: 'https://mcpcoop.org/znichka_icon.webp',
    accent: {
      bar: 'bg-violet-500',
      glow: 'rgba(139,92,246,0.12)',
      border: 'hover:border-violet-400/40',
    },
    iosHref: 'https://apps.apple.com/app/znichka-match-numbers-puzzle/id6793941409',
    androidHref: 'https://play.google.com/store/apps/details?id=org.mcpcoop.znichka',
  },
];
