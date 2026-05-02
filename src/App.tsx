import { AppLayout } from './components/layout/app-layout';
import { ProfileSection } from './components/sections/profile-section';
import { SkillsSection } from './components/sections/skills-section';
import { ExperienceSection } from './components/sections/experience-section';
import { EducationSection } from './components/sections/education-section';
import './config/i18n';

function App() {
  return (
    <AppLayout>
      <div className="space-y-12">
        <ProfileSection />
        <SkillsSection />
        <ExperienceSection />
        <EducationSection />
      </div>
    </AppLayout>
  );
}

export default App;
