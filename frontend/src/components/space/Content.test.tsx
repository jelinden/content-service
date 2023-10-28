import React from 'react';
import { render, screen } from '@testing-library/react';
import Content from './Content';

test('renders content service content page', () => {
  render(<Content />);
  const element = screen.getAllByText(/Content/);
  expect(element[0]).toBeInTheDocument();
});