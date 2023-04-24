# Event Sourcing

    - Planning Events that run through the system

## Explore Business Domain

    - Identify Domain Events
        - The first step is to begin identifying events. 
            - Events are “things that happen” in a process or system. 
        - Important events trigger reactions, so understanding the causal event can help you understand how the system operates and why. 
        - Keep in mind that events are always noted in the past tense. 
        - As your team adds events to the canvas, start organizing them in sequential order over time.
    - Connect Domain events to Commands
        - With your events outlined, you can begin to evaluate each event for its causes and consequences. 
            - In other words, ask yourself what triggered this event (e.g., users, other events, or external systems). 
        - The trigger or cause of the event is noted as a command. 
        - Commands are traditionally documented on blue sticky notes in the present tense and often represent user interactions with the system (e.g., “Submit a purchase order”). 
            - However, your team may want to document commands as both user and system actions.
    - Connect events to reactions
        - The other side of the event process is what happens as a result of the event. 
        - These are called reactions. Reactions are the necessary actions or results following an event and are noted in the present tense. 
            - For example, an event-reaction process flow might state, “When a new account is created, we will send a confirmation email.” 
        - Together, commands, reactions, and events will successfully map out the complete cause-and-effect processes of the system.

## Combine with Domain Driven Design

    - For high-level event storming, the process can end once your team has added domain events, commands, and reactions. However, event storming can be combined with the technique of domain-driven design to define the structure of your system and send your team on the way to implementation. To put it simply, you can start grouping together modules with an element called bounded contexts. Draw a box or circle around the related modules—or in Lucidchart, you can use containers—and label them. Then you can use arrows to begin context mapping—in other words, to show how the modules within a bounded context interact with other contexts. For example, you might have grouped together commands, reactions, and events that deal with payment. The payment context is linked to the shipping context, because once a payment is processed, the system needs to fire a command to ship the product that has been ordered.