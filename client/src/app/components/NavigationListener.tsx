'use client';

import { useEffect } from 'react';
import { usePathname } from 'next/navigation';
import { useDispatch } from 'react-redux';
import { setCurrentPage } from '../core/navigationSlice';

export const NavigationListener = () => {
  const pathname = usePathname();
  const dispatch = useDispatch();

  const getLastPath = (path: string) => {
    const cleanPath = path.replace(/^\/|\/$/g, '').split('?')[0];
    const segments = cleanPath.split('/');
    return segments[segments.length - 1] || '';
  }

  useEffect(() => {
    if (!pathname) return;
    
    const lastSegment = getLastPath(pathname)

    dispatch(setCurrentPage(lastSegment));
  }, [pathname, dispatch]);

  return null;
};