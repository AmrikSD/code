import Countdown from "@/app/Countdown"
import React from "react"

export default function Start() {
    return <>
        <h2>We are getting married!</h2>
        <Countdown target={new Date("2025-02-02")} />
        <p>You&apos;re invited!</p>
        <p>Scroll down for some <a href="#info">more info</a>!</p>
        <p><a href="#location">location</a></p>
        <p><a href="#rsvp">rsvp</a></p>
        <p><a href="#faq">faq</a></p>
    </>
}
