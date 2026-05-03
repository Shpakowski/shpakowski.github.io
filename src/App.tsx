import { AppLayout } from './components/layout/app-layout';
import { HeroSection } from './components/sections/hero-section';
import { SkillsSection } from './components/sections/skills-section';
import { ExperienceSection } from './components/sections/experience-section';
import { useTranslation } from 'react-i18next';
import './config/i18n';

function App() {
  const { t } = useTranslation();

  return (
    <AppLayout>
      <div className="space-y-12">
        <h1 className="text-xl md:text-2xl lg:text-3xl font-bold tracking-tight text-center mb-6">
          {t('profile.headline')}
        </h1>
        
        <HeroSection />
        <SkillsSection />
        <ExperienceSection />
      </div>
    </AppLayout>
  );
}

export default App;
