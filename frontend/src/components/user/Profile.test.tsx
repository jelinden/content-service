import React from 'react';
import { render, screen } from '@testing-library/react';
import Profile from './Profile';

test('renders content service profile page', () => {
  render(<Profile />);
  const element = screen.getAllByText(/Profile/);
  expect(element[0]).toBeInTheDocument();
});