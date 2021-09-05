import React from 'react';
import { render, screen } from '@testing-library/react';
import TreeComponent from './tree';

test('renders learn react link', () => {
  render(<TreeComponent />);
  const linkElement = screen.getByText(/learn react/i);
  expect(linkElement).toBeInTheDocument();
});
