# Events

## Domain Events

    - Identify Domain Events
        - The first step is to begin identifying events. 
            - Events are “things that happen” in a process or system. 
        - Important events trigger reactions, so understanding the causal event can help you understand how the system operates and why. 
        - Keep in mind that events are always noted in the past tense. 
        - As your team adds events to the canvas, start organizing them in sequential order over time.

    - Domain events are also called business events. An event is an action that occurred in the system at a specific time.
        - The first step in the event-storming process consists of these actions:
            - Identifying all the relevant events in the domain and the specific process that you're analyzing
            - Writing a short description of each event on a sticky note
            - Placing all the event sticky notes in sequence on a timeline
        - The act of writing event descriptions often results in questions that need to be resolved later or discussions about definitions that need to be recorded to ensure that everyone agrees on basic domain concepts.
        - Verb Past Tense
        - relevant to domain
    - External systems produce events.
    
    - Events can be created by commands or by external systems, including Internet of Things (IoT) devices. 
        - Events can be triggered by the processing of other events or by a period of elapsed time. 
    - When an event is repeated or occurs regularly on a schedule, draw a clock or calendar icon in the corner of the sticky note for that event.
        - As you identify events and place them into a timeline, you might find independent subsequent events that aren't directly coupled to each other and that represent different perspectives of the system, but that occur in overlapped periods of time. 
        - You can address those parallel event streams by putting them into separate swimlanes that are identified by horizontal blue painter's tape.
        - As you organize the events into a timeline, possibly with swimlanes, you can identify pivotal events. 
        - Pivotal events indicate major changes in the domain and often form the boundary between one phase of the system and another. 
        - They typically separate, or are a bounded context in Domain-Driven Design terms. 
        - Pivotal events are identified by vertical blue painter’s tape that crosses all the swimlanes.
