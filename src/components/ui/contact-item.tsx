import type { ReactNode } from 'react';

interface ContactItemProps {
  icon: ReactNode;
  text: ReactNode;
  href?: string;
}

export function ContactItem({ icon, text, href }: ContactItemProps) {
  if (!text) return null;

  if (href) {
    return (
      <a
        href={href}
        target="_blank"
        rel="noopener noreferrer"
        className="flex items-start gap-3.5 text-muted-foreground hover:text-primary transition-colors group"
      >
        <span className="text-primary group-hover:text-primary transition-colors mt-0.5 shrink-0">{icon}</span>
        <span className="break-words leading-relaxed">{text}</span>
      </a>
    );
  }

  return (
    <div className="flex items-start gap-3.5 text-muted-foreground">
      <span className="text-primary mt-0.5 shrink-0">{icon}</span>
      <span className="break-words leading-relaxed">{text}</span>
    </div>
  );
}
