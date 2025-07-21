'use client';

import { useStore } from 'react-redux';
import { useEffect } from 'react';
import { NavigationListener } from '../components/NavigationListener';

export const NavigationProvider = ({ children }: { children: React.ReactNode }) => {
  const store = useStore();

  useEffect(() => {
  }, [store]);

  return (
    <>
      <NavigationListener />
      {children}
    </>
  );
};