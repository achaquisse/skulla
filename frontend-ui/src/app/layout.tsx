import './globals.css'
import {Inter} from 'next/font/google'
import Image from "next/image";

const inter = Inter({subsets: ['latin']})

export const metadata = {
    title: 'Omniscience School',
    description: 'School Management App',
}

export default function RootLayout({
                                       children,
                                   }: {
    children: React.ReactNode
}) {
    return (
        <html lang="en">
        <body>
        <nav className="navbar bg-body-tertiary">
            <div className="container">
                <a className="navbar-brand" href="#">
                    <Image src="/omniscience-school.png" alt="Logo" width="125" height="26"
                           className="d-inline-block align-text-top"/>
                </a>
            </div>
        </nav>
        <div className="container">
            {children}
        </div>
        </body>
        </html>
    )
}
