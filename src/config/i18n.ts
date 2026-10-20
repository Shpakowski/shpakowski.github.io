import i18n from 'i18next';
import { initReactI18next } from 'react-i18next';
import LanguageDetector from 'i18next-browser-languagedetector';

import en from '@/locales/en/translation.json';
import ru from '@/locales/ru/translation.json';

export const defaultNS = 'translation';
export const resources = {
  en: {
    translation: en,
  },
  ru: {
    translation: ru,
  },
} as const;

const SEO_META = {
  en: {
    title: 'Vitali Shpakowski — Principal Software Engineer | Agentic AI Engineer',
    description:
      'Vitali Shpakowski — Principal Software Engineer & Agentic AI Engineer with 14+ years of experience in System Design, Full-Stack Web (React, Next.js, Angular, Node.js, NestJS), Cross-Platform Mobile (Flutter, iOS/Android), and Agentic AI Systems (Harness, Context & Graph Engineering, MCP).',
  },
  ru: {
    title: 'Виталий Шпаковский — Principal Software Engineer | Agentic AI Engineer',
    description:
      'Виталий Шпаковский (Vitali Shpakowski) — Principal Software Engineer и Agentic AI Engineer с 14+ годами опыта в System Design, Full-Stack разработке (React, Next.js, Angular, Node.js, NestJS), кроссплатформенном мобайле (Flutter, iOS/Android) и создании агентных ИИ-систем (Harness, Context & Graph Engineering, MCP).',
  },
} as const;

function syncDocumentSeo(lng?: string) {
  if (typeof document === 'undefined') return;
  const normalized = lng?.startsWith('ru') ? 'ru' : 'en';
  document.documentElement.lang = normalized;
  document.title = SEO_META[normalized].title;
  const descMeta = document.querySelector('meta[name="description"]');
  if (descMeta) {
    descMeta.setAttribute('content', SEO_META[normalized].description);
  }
}

i18n
  .use(LanguageDetector)
  .use(initReactI18next)
  .init({
    resources,
    fallbackLng: 'en',
    defaultNS,
    interpolation: {
      escapeValue: false, // react already safes from xss
    },
  })
  .then(() => {
    syncDocumentSeo(i18n.resolvedLanguage || i18n.language);
  });

i18n.on('languageChanged', (lng) => {
  syncDocumentSeo(lng);
});

export default i18n;
