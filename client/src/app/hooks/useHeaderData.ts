import { useAppSelector } from "./hook";


export const useHeaderData = () => {
    return useAppSelector((state) => state.navigation.header)
}