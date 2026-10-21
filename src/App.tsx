import { AppLayout } from './components/layout/app-layout';
import React, { Suspense } from 'react';

const HeroSection = React.lazy(() => import('./components/sections/hero-section').then(m => ({ default: m.HeroSection })));
const SkillsSection = React.lazy(() => import('./components/sections/skills-section').then(m => ({ default: m.SkillsSection })));
const ProjectsSection = React.lazy(() => import('./components/sections/projects-section').then(m => ({ default: m.ProjectsSection })));
const ExperienceSection = React.lazy(() => import('./components/sections/experience-section').then(m => ({ default: m.ExperienceSection })));
const EducationSection = React.lazy(() => import('./components/sections/education-section').then(m => ({ default: m.EducationSection })));
import { LayoutProvider } from './context/LayoutContext';
import { useTranslation } from 'react-i18next';
import './config/i18n';

// 🎓 React Trend: Реализовано Code Splitting для ускорения FCP (First Contentful Paint).
function App() {
  const { t } = useTranslation();

  return (
    <LayoutProvider>
      <AppLayout>
        <div className="space-y-12">
        <h1 className="text-xl md:text-2xl lg:text-3xl font-light tracking-tight text-center text-primary mb-6">
          {t('profile.headline')}
        </h1>
        
        <Suspense fallback={<div className="h-32 flex items-center justify-center text-muted">{t('labels.loading')}</div>}>
          <HeroSection />
          <SkillsSection />
          <ProjectsSection />
          <ExperienceSection />
          <EducationSection />
        </Suspense>
        </div>
      </AppLayout>
    </LayoutProvider>
  );
}

export default App;
