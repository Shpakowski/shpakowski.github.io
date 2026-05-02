import * as React from 'react';
import { cn } from '../../utils/cn';

interface SectionBlockProps extends React.HTMLAttributes<HTMLElement> {
  id?: string;
  title: string;
  children: React.ReactNode;
}

export function SectionBlock({ id, title, children, className, ...props }: SectionBlockProps) {
  return (
    <section id={id} className={cn("mb-12 scroll-mt-12", className)} {...props}>
      <h2 
        className="mb-6 text-3xl font-normal uppercase tracking-wide text-primary"
        style={{
          WebkitTextStroke: '1.5px #000',
          paintOrder: 'stroke fill',
          textShadow: '0px 4px 10px rgba(0, 0, 0, 0.25)'
        }}
      >
        {title}
      </h2>
      <div className="w-full">
        {children}
      </div>
    </section>
  );
}
