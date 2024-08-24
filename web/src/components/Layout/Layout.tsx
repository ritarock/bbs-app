import { ReactNode } from "react";
import Header from "../Header";

interface LayoutProps {
  children: ReactNode
}

const Layout: React.FC<LayoutProps> = ({ children }) => {
  return (
    <>
      <Header />
      <hr />
      <br />
      {children}
    </>
  )
}

export default Layout
