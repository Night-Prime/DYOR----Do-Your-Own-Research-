import {z} from 'zod'

export const SignupFormSchema = z.object({
    first_name: z.string().min(1, "First name is required"),
    last_name: z.string().min(1, "Last name is required"),
    avatar: z.string().optional(),
    email: z.string().email("Invalid email address"),
    role: z.literal('user'),
    password: z.string().min(6, "Password must be at least 6 characters long"),
});

export type FormState = {
    errors?: {
        first_name?: string[];
        last_name?: string[];
        avatar?: string[];
        email?: string[];
        role?: string[];
        password?: string[];
    };
};

export const LoginFormSchema = z.object({
    email: z.string().email('Invalid email address'),
    password: z.string().min(8, 'Password must be at least 8 characters')
});


export type loginFormState = {
    errors?: {
        email?: string[];
        password?: string[];
    };
};
