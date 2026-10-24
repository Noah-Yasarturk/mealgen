import HamburgerMenu from "../navMenu/hamburgerMenu.tsx";
import "../header/header.scss"

type HeaderProps = {
    pageName: string;
}

export default function AppHeader({pageName}: HeaderProps) {
    return (
        <div className="appHeader">
            <HamburgerMenu style={{ flex: 1 }} />
            <h2 className="pageName">{pageName}</h2>
        </div>

    )
}