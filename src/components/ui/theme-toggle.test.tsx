import { describe, it, expect, beforeEach } from 'vitest';
import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { ThemeToggle } from './theme-toggle';

describe('ThemeToggle', () => {
  beforeEach(() => {
    localStorage.clear();
    document.documentElement.className = '';
  });

  it('renders default light theme correctly', () => {
    render(<ThemeToggle />);
    expect(document.documentElement).not.toHaveClass('dark');
  });

  it('toggles theme on click and updates localStorage and DOM', async () => {
    const user = userEvent.setup();
    render(<ThemeToggle />);
    
    const button = screen.getByRole('button', { name: /toggle theme/i });
    
    // Switch to dark
    await user.click(button);
    expect(localStorage.getItem('theme')).toBe('dark');
    expect(document.documentElement).toHaveClass('dark');
    
    // Switch to light
    await user.click(button);
    expect(localStorage.getItem('theme')).toBe('light');
    expect(document.documentElement).not.toHaveClass('dark');
  });
});
