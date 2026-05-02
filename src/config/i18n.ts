import i18n from 'i18next';
import { initReactI18next } from 'react-i18next';
import LanguageDetector from 'i18next-browser-languagedetector';

import en from '../locales/en/translation.json';
import ru from '../locales/ru/translation.json';

export const defaultNS = 'translation';
export const resources = {
  en: {
    translation: en,
  },
  ru: {
    translation: ru,
  },
} as const;

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
  });

export default i18n;
