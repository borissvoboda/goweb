CREATE TABLE IF NOT EXISTS clients (
     id              INTEGER PRIMARY KEY AUTOINCREMENT,
     company_name        TEXT NOT NULL UNIQUE,
     uid          TEXT NOT NULL UNIQUE,
     email   TEXT NOT NULL,
     note TEXT
 );

 INSERT OR IGNORE INTO clients (company_name, uid, email, note)
     VALUES
        ('TechNova Solutions',       'CHE-112.345.678',        'info@technova.ch',    'Main supplier - cloud services'),         
        ('GreenLeaf Logistics',      'CHE-987.654.321',        'contact@greenleaf.eu','Prefers phone communication'),
        ('PixelForge Studio',        'CHE-456.789.123',        'hello@ pixelforge.com', NULL),
        ('BioFarm Organics AG',      'CHE-321.654.987',        'office@biofarm.ch',   'Seasonal delivery schedule'),
        ('Sumatra CRM',              'sumatra-crm',            'sales@sumatra.com',           'Modern CRM focused on simplicity'),
        ('Freshsales (Freshworks)',  'freshsales',             'sales@freshworks.com',        'Part of Freshworks suite'),
        ('Pipedrive',                'pipedrive',              'support@pipedrive.com',       'Sales pipeline focused CRM'),
        ('Monday.com CRM',           'monday-crm',             'info@monday.com',             'Highly customizable work OS'),
        ('ActiveCampaign',           'activecampaign',         'sales@activecampaign.com',    'Strong in email marketing automation'),
        ('Less Annoying CRM',        'less-annoying-crm',      'support@lessannoyingcrm.com', 'Very simple and affordable'),
        ('Agile CRM',                'agile-crm',              'sales@agilecrm.com',          'Good free plan available'),
        ('Bitrix24',                 'bitrix24',               'sales@bitrix24.com',          'All-in-one platform with many tools'),
        ('Creatio',                  'creatio',                'info@creatio.com',            'Low-code process automation'),
        ('Apptivo',                  'apptivo',                'support@apptivo.com',         NULL);


SELECT * FROM clients;