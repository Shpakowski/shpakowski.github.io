import { describe, it, expect, vi } from 'vitest';
import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { LanguageSwitch } from './language-switch';

const changeLanguageMock = vi.fn();
vi.mock('react-i18next', async (importOriginal) => {
  const actual = await importOriginal<typeof import('react-i18next')>();
  return {
    ...actual,
    useTranslation: () => ({
      i18n: {
        language: 'en',
        changeLanguage: changeLanguageMock,
      },
    }),
  };
});

describe('LanguageSwitch', () => {
  it('renders correctly', () => {
    render(<LanguageSwitch />);
    const button = screen.getByRole('button', { name: /toggle language/i });
    expect(button).toBeInTheDocument();
  });

  it('calls changeLanguage on click', async () => {
    const user = userEvent.setup();
    render(<LanguageSwitch />);
    
    const button = screen.getByRole('button', { name: /toggle language/i });
    await user.click(button);
    
    expect(changeLanguageMock).toHaveBeenCalledWith('ru');
  });
});
