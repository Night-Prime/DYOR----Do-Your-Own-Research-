import { useEffect } from "react"
import { useRouter } from "next/navigation"
import { useAppSelector } from "../hooks/hook"

const ProtectedRoute = ({children}: {children : React.ReactNode}) => {
    const {isAuthenticated} = useAppSelector(state => state.auth)
    const router = useRouter()

    useEffect(() => {
        if(!isAuthenticated) {
            // need to setup an action that triggers a modal
            router.push('/')
        }
    }, [isAuthenticated, router])

    return isAuthenticated ? <>{children}</> : null;
}

export default ProtectedRoute;