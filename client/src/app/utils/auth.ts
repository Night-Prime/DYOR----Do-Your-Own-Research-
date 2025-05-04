"use client"

import { SignupFormSchema, FormState, LoginFormSchema, loginFormState } from "./validation";
import axios from "axios";

export async function signup(state: FormState, formData: FormData) {
    // First, Validation
    const validatedData = SignupFormSchema.safeParse({
        first_name: formData.get('first_name'),
        last_name: formData.get('last_name'),
        avatar: formData.get('avatar'),
        email: formData.get('email'),
        password: formData.get('password'),
        role: formData.get('role')
    })

    if (!validatedData.success && validatedData.error) {
        return {
            errors: validatedData.error.flatten().fieldErrors
        }
    }

    // Then make the api call
    try {
        const response = await axios.post(`${process.env.NEXT_PUBLIC_API_URL}/user/signup`, validatedData.data, {
            headers: {
                'Content-Type': 'application/json',
            },
        });
        console.log("Response: ", response)
        return {
            success: true,
            data: response.data,
        };
    } catch (error) {
        if (axios.isAxiosError(error) && error.response) {
            return {
                errors: error.response.data.errors || { general: ['An error occurred during signup.'] },
            };
        }
        return {
            errors: { general: ['Failed to connect to the signup service.'] },
        };
    }
}

export async function login(state: loginFormState, formData: FormData) {
    // Validate form data
    const validatedData = LoginFormSchema.safeParse({
        email: formData.get('email'),
        password: formData.get('password'),
    });

    if (!validatedData.success && validatedData.error) {
        return {
            success: false,
            errors: validatedData.error.flatten().fieldErrors
        };
    }

    try {
        const response = await axios.post(
            `${process.env.NEXT_PUBLIC_API_URL}/user/login`,
            {
                email: validatedData.data.email,
                password: validatedData.data.password
            },
            {
                headers: {
                    'Content-Type': 'application/json',
                },
                withCredentials: true,
            }
        );
        console.log("Response: ", response)


        return {
            success: true,
            data: response.data,
        };
    } catch (error) {
        console.log("Err: ", error);
        if (axios.isAxiosError(error) && error.response) {
            return {
                success: false,
                errors: error.response.data.errors || { 
                    general: ['An error occurred during login.'] 
                },
            };
        }
        return {
            success: false,
            errors: { general: ['Failed to connect to the login service.'] },
        };
    }
}