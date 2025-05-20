import Dashbar from "../components/Dashbar";
import Sidebar from "../components/Sidebar";
import { AuthProvider } from "../shared/AuthProvider";

export default function DashboardLayout({
    children,
}: Readonly<{
    children: React.ReactNode;
}>) {
    return (
        <AuthProvider>
            <section className="w-screen h-screen overflow-hidden bg-white">
                <main className="w-full h-full flex-grow grid grid-cols-[0.125fr_2fr] text-lime-800 p-2 gap-2">
                    <div className="bg-gray-200 rounded-3xl">
                        <Sidebar />
                    </div>
                    <div className="w-full h-full flex flex-col bg-gray-100 rounded-3xl overflow-hidden">
                        <div className="h-[10%] min-h-24 bg-gray-100">
                            <Dashbar />
                        </div>

                        <div className="flex-1 overflow-auto z-10">
                            <div className="h-full flex flex-col justify-center items-center">
                                {children}
                            </div>
                        </div>
                    </div>
                </main>
            </section>
        </AuthProvider>
    )
}