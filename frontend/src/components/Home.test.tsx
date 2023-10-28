import React from 'react';
import { render, screen } from '@testing-library/react';
import Home from './Home';

test('renders content service home page', () => {
  render(<Home />);
  const element = screen.getAllByText(/Content service is a/);
  const elementImg = screen.getAllByAltText(/content/);
  expect(element[0]).toBeInTheDocument();
  expect(elementImg[0]).toBeVisible();
});