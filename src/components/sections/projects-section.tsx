import { useTranslation } from 'react-i18next';
import { motion } from 'framer-motion';
import { PROJECTS_CONFIG } from '@/config';
import { SectionHeader } from '@/components/ui';

export function ProjectsSection() {
  const { t } = useTranslation();

  return (
    <section id="projects-section" className="py-2">
      <SectionHeader title={t('sections.projects')} />

      <div className="grid grid-cols-1 sm:grid-cols-3 gap-4">
        {PROJECTS_CONFIG.map((project, i) => (
          <motion.a
            key={project.id}
            href={project.href || undefined}
            target={project.href ? '_blank' : undefined}
            rel={project.href ? 'noopener noreferrer' : undefined}
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ delay: 0.1 + i * 0.12, duration: 0.4 }}
            whileHover={{ y: -6, transition: { duration: 0.2, ease: "easeOut" } }}
            className={`
              group relative flex flex-col overflow-hidden rounded-xl cursor-pointer
              bg-card border border-border
              ${project.accent.border}
              transition-colors duration-200
            `}
            style={{
              boxShadow: '0 1px 4px rgba(0,0,0,0.06), 0 2px 8px rgba(0,0,0,0.04)',
            }}
            onMouseEnter={e => {
              (e.currentTarget as HTMLElement).style.boxShadow =
                `0 8px 32px ${project.accent.glow}, 0 2px 8px rgba(0,0,0,0.06)`;
            }}
            onMouseLeave={e => {
              (e.currentTarget as HTMLElement).style.boxShadow =
                '0 1px 4px rgba(0,0,0,0.06), 0 2px 8px rgba(0,0,0,0.04)';
            }}
          >
            {/* Accent top bar */}
            <div
              className={`h-[3px] w-full ${project.accent.bar} origin-left scale-x-[0.35] group-hover:scale-x-100 transition-transform duration-500 ease-out`}
            />

            {/* Logo area — background-image so any size fits */}
            <div
              className="mx-auto mt-6 h-16 w-16 rounded-2xl bg-center bg-cover bg-no-repeat transition-transform duration-300 group-hover:scale-110"
              style={{ backgroundImage: `url(${project.logo})` }}
              aria-hidden="true"
            />

            {/* Content */}
            <div className="flex flex-col items-center text-center px-6 pt-4 pb-6 gap-2">
              {/* Name — uppercase, IBM Plex Sans, thin/light, foreground */}
              <h3 className="text-sm font-normal tracking-[0.12em] uppercase text-foreground">
                {t(project.nameKey)}
              </h3>

              {/* Description — tight, muted */}
              <p className="text-sm text-left w-full leading-snug text-muted">
                {t(project.descriptionKey)}
              </p>
            </div>
          </motion.a>
        ))}
      </div>
    </section>
  );
}
