import Section from "@/app/section"
import React from "react"
import Start from "./start"
import Location from "./location"
import Faq from "./faq"
import Rsvp from "./rsvp"

import img from "../../public/image (1).jpg"

export default function Index() {
    return <>
        <Section name="start"> 
            <Start />
        </Section>
        <Section name="location"> 
            <Location />
        </Section>
        <Section name="faq"> 
            <Faq />
        </Section>
        <Section name="rsvp"> 
            <Rsvp />
        </Section>
        <Section name="kate"> 
            <img src={img.src} alt="kate and amrik"/>
            <h2>Made with ❤️ and 🫖 and 🍫 by 🥚 and 🍅</h2>
        </Section>
    </>
}
