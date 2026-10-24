import { PiHamburgerLight } from "react-icons/pi";
import "../navMenu/hamburgerMenu.scss";


export default function HamburgerMenu({ style }) {
    return (
        <div id="piHamburgerIconContainer" style={{...style}}>
            <PiHamburgerLight />
        </div>
    )
}