'use client';

import React from 'react';
import { Button } from '@/components/ui/button';
import { Download, Share, Plus } from 'lucide-react';
import { toast } from 'sonner';
import { useInstallPWA } from './hook';
import { cn } from '@/lib/utils';

interface InstallButtonProps {
  className?: string;
  variant?: 'default' | 'outline' | 'secondary' | 'ghost' | 'link';
  size?: 'default' | 'sm' | 'lg' | 'icon';
  children?: React.ReactNode;
}

export function InstallButton({
  className = '',
  variant = 'default',
  size = 'default',
  children,
}: InstallButtonProps) {
  const { installPWA, isInstallable, isIOS, isInstalled } = useInstallPWA();

  if (isInstalled) return null;

  const handleInstall = async () => {
    const result = await installPWA();

    if (result.isIOS) {
      toast.info(
        <div className='flex flex-col gap-2'>
          <h3 className='font-medium'>Install to Home Screen</h3>
          <p className='text-sm flex items-center gap-1'>
            Tap the share button <Share className='h-4 w-4' />
            and then &quot;Add to Home Screen&quot; <Plus className='h-4 w-4' />
          </p>
        </div>,
        { duration: 10000 }
      );
      return;
    }

    if (result.outcome === 'accepted') {
      toast.success('Installing app');
    }
  };

  return (
    <Button
      className={cn('gap-2', className)}
      variant={variant}
      size={size}
      onClick={handleInstall}
      disabled={!isInstallable && !isIOS}
    >
      <Download className='h-4 w-4' />
      {children || 'Install App'}
    </Button>
  );
}
