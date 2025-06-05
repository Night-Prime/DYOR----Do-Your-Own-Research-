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

export const UpdatePortfolioSchema = z.object({
    portfolio: z.object({
        id: z.string().uuid(),
        user_id: z.string().uuid(),
        name: z.string().min(1).optional(),
        total_value: z.number().nonnegative().optional(),
        investment_goals: z.array(z.string().min(1)).optional(),
        asset_preference: z.array(z.string().min(1)).optional(),
    }),
    assets: z.object({
        add: z.array(
            z.object({
                symbol: z.string().min(1),
                name: z.string().min(1),
                type: z.enum(["stock", "crypto"]),
                quantity: z.number().nonnegative().optional(),
                current_price: z.number().nonnegative().optional(),
                volume: z.number().nonnegative().optional(),
            })
    ).optional(),
    // update: z.array(
    //         z.object({
    //             id: z.string().uuid(),
    //             symbol: z.string().min(1),
    //             name: z.string().min(1).optional(),
    //             quantity: z.number().nonnegative().optional(),
    //             current_price: z.number().nonnegative().optional(),
    //             volume: z.number().nonnegative().optional(),
    //         })
    //     ).optional(),
    // delete: z.array(z.string().uuid()).optional(),
}).optional(),
});

export type UpdatePortfolioFormState = z.infer<typeof UpdatePortfolioSchema>;