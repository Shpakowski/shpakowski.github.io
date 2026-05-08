export function AchievementItem({ text }: { text: string }) {
  // Parse **bold** markdown
  const parts = text.split(/\*\*(.*?)\*\*/g);
  return (
    <li className="text-sm text-muted leading-relaxed pl-4 relative before:content-[''] before:absolute before:left-0 before:top-[9px] before:w-1.5 before:h-1.5 before:rounded-full before:bg-timeline-dot/40">
      {parts.map((part, i) =>
        i % 2 === 1 ? (
          <strong key={i} className="font-semibold text-foreground">{part}</strong>
        ) : (
          <span key={i}>{part}</span>
        )
      )}
    </li>
  );
}
