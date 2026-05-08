import { describe, it, expect } from 'vitest';
import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { LayoutProvider, useLayout } from './LayoutContext';

const TestComponent = () => {
  const { activeExperienceIndex, forceExpandSignal, openExperience } = useLayout();
  return (
    <div>
      <span data-testid="index">{activeExperienceIndex === null ? 'null' : activeExperienceIndex}</span>
      <span data-testid="signal">{forceExpandSignal}</span>
      <button onClick={() => openExperience(2, true)}>Open with expand</button>
      <button onClick={() => openExperience(null, false)}>Close</button>
    </div>
  );
};

describe('LayoutContext', () => {
  it('provides default values and updates state', async () => {
    const user = userEvent.setup();
    render(
      <LayoutProvider>
        <TestComponent />
      </LayoutProvider>
    );

    expect(screen.getByTestId('index')).toHaveTextContent('null');
    expect(screen.getByTestId('signal')).toHaveTextContent('0');

    await user.click(screen.getByText('Open with expand'));

    expect(screen.getByTestId('index')).toHaveTextContent('2');
    expect(screen.getByTestId('signal')).toHaveTextContent('1');

    await user.click(screen.getByText('Close'));

    expect(screen.getByTestId('index')).toHaveTextContent('null');
    expect(screen.getByTestId('signal')).toHaveTextContent('0');
  });
});
