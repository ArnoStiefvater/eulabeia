# Copyright (C) 2021 Greenbone Networks GmbH
#
# SPDX-License-Identifier: GPL-2.0-or-later
#
# This program is free software; you can redistribute it and/or
# modify it under the terms of the GNU General Public License
# as published by the Free Software Foundation; either version 2
# of the License, or (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with this program; if not, write to the Free Software
# Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA 02110-1301 USA.

# see /opt/local/share/cmake-3.16/Modules/FindGTest.cmake
find_package(cgreen 1.1.0)
# Now simply link against gtest as needed. Eg
add_executable(all_tests src/all.c src/start_scan.c src/scan_porgress.c)
target_link_libraries(all_tests ${CGREEN_LIBRARIES} Eulabeia::eulabeia)


add_test(NAME test_start_scan_fail
    COMMAND all_tests start_scan_returns_error_on_publish_fail)
add_test(NAME test_start_scan_success
    COMMAND all_tests start_scan_success)
add_test(NAME test_progress_fail
    COMMAND all_tests scan_progress_failures)
