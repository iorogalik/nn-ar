CREATE TABLE IF NOT EXISTS public.devices
(
    id                serial PRIMARY KEY,
    organization_id   integer NOT NULL REFERENCES public.organizations(id),
    room_id           integer NOT NULL REFERENCES public.rooms(id),
    guid              uuid NOT NULL UNIQUE,
    inventory_number  varchar(255) NOT NULL UNIQUE,
    serial_number     varchar(255) NOT NULL UNIQUE,
    characteristics   varchar(500) NOT NULL,
    category          varchar(255) NOT NULL,
    units             varchar(50) NOT NULL,
    powerConsumption  integer NOT NULL,
    created_date      timestamptz NOT NULL,
    updated_date      timestamptz NOT NULL,
    deleted_date      timestamptz
);
