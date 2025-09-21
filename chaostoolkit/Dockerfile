FROM python:3.11-slim

# OS packages minimal
RUN apt-get update && apt-get install -y --no-install-recommends \
    curl ca-certificates bash jq libcairo2 libcairo2-dev pandoc \
    texlive-latex-base texlive-xetex texlive-fonts-recommended texlive-latex-recommended lmodern \
 && rm -rf /var/lib/apt/lists/*

# Install Chaos Toolkit + plugin yang umum dipakai
# (requests untuk HTTP checks; toxiproxy untuk integrasi Toxiproxy)
RUN pip install --no-cache-dir \
    chaostoolkit \
    chaostoolkit-toxiproxy \
    chaostoolkit-reporting

# User non-root
RUN useradd -ms /bin/bash chaos
USER chaos

WORKDIR /experiments
