// GlobalAlert.tsx
import { hideAlert } from '../core/alertSlice';
import { useAppSelector, useAppDispatch } from '../hooks/hook';
import { DyorAlert } from './Alert';

export const GlobalAlert = () => {
  const alert = useAppSelector((state) => state.alert);
  const dispatch = useAppDispatch();
    console.log("triggered!");
  return (
    <DyorAlert
      type={alert.type}
      message={alert.message}
      open={alert.open}
      autoClose={alert.autoClose}
      autoCloseDuration={alert.autoCloseDuration}
      onClose={() => dispatch(hideAlert())}
    />
  );
};